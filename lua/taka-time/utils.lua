local M = {}
local config = require("taka-time.config")

function M.get_binary_path()
	local plugin_root = vim.fn.fnamemodify(debug.getinfo(1).source:sub(2), ":h:h:h")
	local bin_path = plugin_root .. "/taka-upload"

	local os_name = vim.loop.os_uname().sysname:lower()
	if string.match(os_name, "windows") ~= nil then
		bin_path = bin_path .. ".exe"
	end

	return bin_path
end

function M.get_binary_path_dahboard()
	local plugin_root = vim.fn.fnamemodify(debug.getinfo(1).source:sub(2), ":h:h:h")
	return plugin_root .. "/taka-dashboard"
end

function M.get_version_file_path()
	local plugin_root = vim.fn.fnamemodify(debug.getinfo(1).source:sub(2), ":h:h:h")
	return plugin_root .. "/.version"
end

-- Helper: Read version from disk
function M.get_installed_version()
	local f = io.open(M.get_version_file_path(), "r")
	if not f then
		return nil
	end
	local v = f:read("*all")
	f:close()
	return v:gsub("%s+", "") -- trim whitespace
end

-- Helper: Write version to disk
function M.write_installed_version(v)
	local f = io.open(M.get_version_file_path(), "w")
	if f then
		f:write(v)
		f:close()
	end
end

-- Helper: Detect OS/Arch
local function get_os_info()
	local uname = vim.loop.os_uname()
	local os = uname.sysname:lower()
	local arch = uname.machine:lower()
	if os == "linux" then
		os = "linux"
	elseif os == "darwin" then
		os = "darwin"
	elseif string.match(os, "windows") ~= nil then
		os = "windows"
	else
		return nil, nil
	end
	if arch == "x86_64" then
		arch = "amd64"
	end
	if arch == "aarch64" then
		arch = "arm64"
	end
	return os, arch
end
--------------------------------------------------------------------------------------------
---binary downloads
--------------------------------------------------------------------------------------------
function M.ensure_binaries()
	local upload_bin = M.get_binary_path()
	local dash_bin = M.get_binary_path_dashboard()
	local target_ver = config.options.binary_version
	local current_ver = M.get_installed_version()

	local upload_exists = vim.fn.filereadable(upload_bin) == 1
	local dash_exists = vim.fn.filereadable(dash_bin) == 1

	-- 1. Check if BOTH are installed and up-to-date
	if upload_exists and dash_exists and current_ver == target_ver then
		return
	end

	-- 2. Update status message
	if current_ver and current_ver ~= target_ver then
		print(string.format("[TakaTime] Updating binaries %s -> %s...", current_ver, target_ver))
	else
		print("[TakaTime] Installing binaries (" .. target_ver .. ")...")
	end

	local os_name, arch = get_os_info()
	if not os_name then
		print("[TakaTime] Auto-install not supported for this OS.")
		return
	end

	-- 3. Helper function to download and set permissions for ANY binary
	local function download_binary(bin_name, dest_path)
		if vim.fn.filereadable(dest_path) == 1 then
			os.remove(dest_path)
		end

		-- Assumes your GitHub release assets are named like: taka-upload-linux-amd64
		local url = string.format(
			"https://github.com/Rtarun3606k/TakaTime/releases/download/%s/%s-%s-%s",
			target_ver,
			bin_name,
			os_name,
			arch
		)

		vim.fn.system({ "curl", "-L", "-o", dest_path, url })
		vim.fn.system({ "chmod", "+x", dest_path })
	end

	-- 4. Execute downloads sequentially
	download_binary("taka-upload", upload_bin)
	download_binary("taka-dashboard", dash_bin)

	-- 5. Update the version file so we don't re-download next time
	M.write_installed_version(target_ver)
	print("[TakaTime] Successfully installed all binaries!")
end



-----------------------------------------------------------------------------------------
-----------------------------------------------------------------------------------------

function M.get_os()
	local os_name = vim.loop.os_uname().sysname

	print(os_name, " os name")
	return os_name
end

-- Helper: Ensure path ends with a separator for safe boundary checking
local function ensure_trailing_slash(path)
	-- If it doesn't already end with a slash (or backslash for Windows), add one
	if not path:match("[\\/]$") then
		return path .. "/"
	end
	return path
end

-- Helper: Check if the current directory is inside an ignored repository
function M.is_ignored(current_dir)
	local config = require("taka-time.config")
	local ignore_list = config.options.ignore_repos or {}

	-- Add a slash to the current directory for safe comparison
	local safe_current = ensure_trailing_slash(current_dir)

	for _, ignored_path in ipairs(ignore_list) do
		-- Add a slash to the ignored path as well
		local safe_ignored = ensure_trailing_slash(ignored_path)

		-- Now it strictly checks folder boundaries, not just partial strings!
		if vim.startswith(safe_current, safe_ignored) then
			return true
		end
	end

	return false
end

return M
