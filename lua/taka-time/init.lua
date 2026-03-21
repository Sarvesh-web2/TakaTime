local M = {}
local config = require("taka-time.config")
local utils = require("taka-time.utils")
-- Ensure this file is actually named 'lua/taka-time/core.lua'
local core = require("taka-time.core")
local storage = require("taka-time.storage")

function M.setup(opts)
	-- 1. Setup Config
	config.setup(opts)

	-- 2. AUTOMATICALLY LOAD SECRET
	if config.options.mongo_uri == "" then
		local saved_uri = storage.load_secret()
		if saved_uri then
			config.options.mongo_uri = saved_uri
		end
	end
	----------------------------------------------------------------------------------------
	-- 3. Create Commands

	-- TakaInit command
	vim.api.nvim_create_user_command("TakaInit", function()
		local uri = vim.fn.inputsecret("Enter your Mongo URI: ")
		if uri and uri ~= "" then
			storage.save_secret(uri)
			config.options.mongo_uri = uri
			print("TakaTime: URI Saved!")
		else
			print("TakaTime: No URI Entered")
		end
	end, {})

	----------------------------------------------------------------------------------
	-- TakaStatus command
	vim.api.nvim_create_user_command("TakaStatus", function()
		if config.options.mongo_uri and config.options.mongo_uri ~= "" then
			print("TakaTime is configured and running.")
		else
			print("TakaTime is NOT configured. Run :TakaInit")
		end
	end, {})

	-----------------------------------------------------------------------------------
	-- Load the ignore list into RAM on startup
	config.options.ignore_repos = storage.load_ignore_list() or {}

	-- Command to IGNORE the current directory
	vim.api.nvim_create_user_command("TakaIgnore", function()
		local cwd = vim.fn.getcwd()

		-- Check if it's already in the RAM list
		if vim.tbl_contains(config.options.ignore_repos, cwd) then
			print("TakaTime: " .. cwd .. " is already being ignored.")
			return
		end

		-- Add to RAM (Instant)
		table.insert(config.options.ignore_repos, cwd)

		-- Save to Disk
		storage.save_ignore_list(config.options.ignore_repos)

		print("TakaTime: Ignored tracking for " .. cwd)
	end, {})

	------------------------------------------------------------------------------------------
	-- Command to TRACK (Undo ignore) for the current directory
	vim.api.nvim_create_user_command("TakaTrack", function()
		local cwd = vim.fn.getcwd()
		local updated_list = {}
		local found = false

		-- Rebuild the list, skipping the current directory
		for _, repo in ipairs(config.options.ignore_repos) do
			if repo == cwd then
				found = true
			else
				table.insert(updated_list, repo)
			end
		end

		if not found then
			print("TakaTime: " .. cwd .. " is already being tracked.")
			return
		end

		-- Update RAM (Instant)
		config.options.ignore_repos = updated_list

		-- Save to Disk
		storage.save_ignore_list(config.options.ignore_repos)

		print("TakaTime: Resumed tracking for " .. cwd)
	end, {})

	-- 4. Ensures Binary Exists
	pcall(utils.ensure_binary)

	-- 5. START TRACKING (The Fix)
	-- We delegate all logic to core. This sets up CursorMoved, TextChanged, AND BufWritePost
	core.setup_listeners()

	-- Start the background sync timer
	core.start_timer()

	-- 6. Exit Handler
	-- We still handle Exit explicitly here or inside core.
	-- Since core.on_exit is public, this is fine, OR core.setup_listeners can handle it.
	-- But usually VimLeavePre is safer in init.lua for plugin lifecycle.
	vim.api.nvim_create_autocmd("VimLeavePre", {
		group = vim.api.nvim_create_augroup("TakaTimeExit", { clear = true }),
		callback = core.on_exit,
	})
end

return M
