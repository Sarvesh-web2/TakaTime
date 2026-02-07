local M = {}

M.defaults = {
	binary_version = "v2.1.0",

	mongo_uri = "",

	debounce_seconds = 2,

	debug = false,
}

M.options = {}

function M.setup(options)
	M.options = vim.tbl_deep_extend("force", M.defaults, options or {})
end

return M
