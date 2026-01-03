local M = {}

-- Default settings
M.defaults = {
  mongo_uri = "",
  debug = false,
  debounce_seconds = 2,
  binary_version = "v1.0.5", -- Bump this when you release new Go versions!
}

-- Active configuration (starts as defaults)
M.options = vim.deepcopy(M.defaults)

function M.setup(user_opts)
  M.options = vim.tbl_deep_extend("force", M.defaults, user_opts or {})
end

return M
