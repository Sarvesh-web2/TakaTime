-- init.lua ok my firat plugin
-- redir > nvim_logs.txt | silent messages | redir END
-- lua/taka-time/init.lua
local M = {}

M.config = {
	mongo_uri = "", -- User must set this!
	debug = true, -- Show print messages
}

-- STATE VARIABLES
local last_active_time = os.time()
local pending_duration = 0
local current_job_id = 0

-- HELPER: Get the path to our compiled Go binary
local function get_binary_path()
	-- Finds the folder where this lua file is, then goes up to finding the 'cmd' folder
	local plugin_root = vim.fn.fnamemodify(debug.getinfo(1).source:sub(2), ":h:h:h")
	return plugin_root .. "/cmd/taka-cli"
end

-- CORE LOGIC: Try to upload if the line is clear
local function attempt_upload()
	-- 1. If an upload is already running, wait.
	if current_job_id ~= 0 then
		return
	end

	--core

	-- 2. If we have less than 10 seconds of data, wait (save bandwidth).
	-- (You can lower this to 2 for testing)
	if pending_duration < 2 then
		return
	end

	-- 3. Prepare Command
	local file_path = vim.fn.expand("%:p")
	local project = vim.fn.fnamemodify(vim.fn.getcwd(), ":t")
	local bin = get_binary_path()
	local fileExtension = vim.fn.fnamemodify(file_path, ":e")

	if fileExtension == "" then
		fileExtension = "text"
	end

	-- Capture the duration we are about to send, then reset the counter
	local time_to_send = pending_duration
	pending_duration = 0

	local cmd = {
		bin,
		"-uri",
		M.config.mongo_uri,
		"-project",
		project,
		"-file",
		file_path,
		"-duration",
		tostring(time_to_send),
		"-language",
		fileExtension,
	}

	if M.config.debug then
		print("[Taka] Syncing " .. time_to_send .. "s...")
	end

	-- 4. Start Async Job
	current_job_id = vim.fn.jobstart(cmd, {
		on_exit = function(_, code)
			current_job_id = 0 -- Unlock the mutex
			if code ~= 0 then
				-- If it failed, put the time back into the pending buffer!
				pending_duration = pending_duration + time_to_send
				if M.config.debug then
					print("[Taka] Failed. Retrying later.")
				end
			elseif M.config.debug then
				print("[Taka] Success.")
			end
		end,
	})
end

-- EVENT: User Saved (:w)
local function on_save()
	local now = os.time()
	local session_time = now - last_active_time
	last_active_time = now

	-- Add this session to our "Bucket"
	pending_duration = pending_duration + session_time

	attempt_upload()
end

-- EVENT: User Quitting (VimLeavePre)
local function on_exit()
	local now = os.time()
	pending_duration = pending_duration + (now - last_active_time)

	if pending_duration > 0 then
		print("[Taka] Uploading final " .. pending_duration .. "s... (Please wait)")

		-- BLOCKING call to ensure it finishes before window closes
		local bin = get_binary_path()
		vim.fn.system({
			bin,
			"-uri",
			M.config.mongo_uri,
			"-project",
			"unknown", -- Context might be lost on exit, keeping it simple
			"-file",
			"closing_session",
			"-duration",
			tostring(pending_duration),
		})
	end
end

function M.setup(opts)
	M.config = vim.tbl_deep_extend("force", M.config, opts or {})
	last_active_time = os.time()

	vim.api.nvim_create_autocmd("BufWritePost", { callback = on_save })
	vim.api.nvim_create_autocmd("VimLeavePre", { callback = on_exit })
end

return M
