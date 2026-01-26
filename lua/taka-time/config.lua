local M = {}

M.defaults = {
	inary_version = "v2.0.5",

	ongo_uri = "",

	ebounce_seconds = 2,

	debug = false,
}

M.options = {}

function M.setup(options)
	.options = vim.tbl_deep_extend("force", M.defaults, options or {})
end

return M
