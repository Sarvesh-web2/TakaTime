local M = {}
local config = require("taka-time.config")
local utils = require("taka-time.utils")

-- STATE
local state = {
	last_active = os.time(),
	pending_duration = 0,
	job_id = 0,
}

-- Internal: The actual upload logic
local function attempt_upload()
	-- 1. Checks
	if state.job_id ~= 0 then
		return
	end -- Busy
	if state.pending_duration < (config.options.debounce_seconds or 2) then
		return
	end -- Not enough data

	-- 2. Snapshot data
	local time_to_send = state.pending_duration
	state.pending_duration = 0

	-- 3. Prepare Args
	local file_path = vim.fn.expand("%:p")
	local project = vim.fn.fnamemodify(vim.fn.getcwd(), ":t")
	local ext = vim.fn.fnamemodify(file_path, ":e")
	if ext == "" then
		ext = "text"
	end

	local cmd = {
		utils.get_binary_path(),
		"-uri",
		config.options.mongo_uri, -- This now works because init.lua populated it!
		"-project",
		project,
		"-language",
		ext,
		"-file",
		file_path,
		"-duration",
		tostring(time_to_send),
		"-editor",
		"NeoVim",
	}

	if config.options.debug then
		print("[Taka] Syncing " .. time_to_send .. "s...")
	end

	-- 4. Run Job
	state.job_id = vim.fn.jobstart(cmd, {
		on_exit = function(_, code)
			state.job_id = 0
			if code ~= 0 then
				-- Failure: Put time back in bucket
				state.pending_duration = state.pending_duration + time_to_send
				if config.options.debug then
					print("[Taka] Failed. Retrying.")
				end
			elseif config.options.debug then
				print("[Taka] Success.")
			end
		end,
	})
end

-- Public: Called on :w
function M.on_save()
	local now = os.time()
	state.pending_duration = state.pending_duration + (now - state.last_active)
	state.last_active = now
	attempt_upload()
end

-- Public: Called on Exit
function M.on_exit()
	local now = os.time()
	state.pending_duration = state.pending_duration + (now - state.last_active)

	if state.pending_duration > 0 then
		print("[Taka] Uploading final data...")
		-- Blocking call for exit
		vim.fn.system({
			utils.get_binary_path(),
			"-uri",
			config.options.mongo_uri,
			"-project",
			"unknown",
			"-language",
			"text",
			"-file",
			"closing_session",
			"-duration",
			tostring(state.pending_duration),
		})
	end
end

-- Helper to start the timer loop
function M.start_timer()
	local timer = vim.loop.new_timer()
	timer:start(
		1000,
		60000,
		vim.schedule_wrap(function()
			local now = os.time()
			-- Add time since last check
			if state.last_active then
				state.pending_duration = state.pending_duration + (now - state.last_active)
			end
			state.last_active = now
			attempt_upload()
		end)
	)
end

return M
