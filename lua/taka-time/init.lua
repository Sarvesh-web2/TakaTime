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

	-- 3. Create Commands
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

	vim.api.nvim_create_user_command("TakaStatus", function()
		if config.options.mongo_uri and config.options.mongo_uri ~= "" then
			print("TakaTime is configured and running.")
		else
			print("TakaTime is NOT configured. Run :TakaInit")
		end
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
