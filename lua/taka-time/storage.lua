local M = {}
local home = vim.fn.expand("~")

-- 1. Create the ~/.takatime folder if it doesn't exist yet!
local taka_dir = home .. "/.takatime"
if vim.fn.isdirectory(taka_dir) == 0 then
    vim.fn.mkdir(taka_dir, "p")
end

-- 2. Define the file paths (Fixed the variable name to ignore_path)
local data_path = vim.fn.stdpath("data") .. "/taka_data.json"
local ignore_path = taka_dir .. "/taka_ignore.json"

function M.save_secret(uri)
    local file = io.open(data_path, "w")
    if file then
        file:write(vim.json.encode({ mongo_uri = uri }))
        file:close()
        print("TakaTime: Mongo URI saved successfully!")
    else
        print("TakaTime Error: Could not save secret to " .. data_path)
    end
end

function M.load_secret()
    local file = io.open(data_path, "r")
    if file then
        local content = file:read("*a")
        file:close()
        -- Safely decode JSON
        local ok, data = pcall(vim.json.decode, content)
        if ok and data.mongo_uri then
            return data.mongo_uri
        end
    end
    return nil
end

function M.save_ignore_list(ignore_list)
    local file = io.open(ignore_path, "w")
    if file then
        -- Encode the Lua table into a JSON array and write it
        file:write(vim.json.encode(ignore_list))
        file:close()
    else
        print("TakaTime Error: Could not save ignore list to " .. ignore_path)
    end
end

function M.load_ignore_list()
    local file = io.open(ignore_path, "r")
    if file then
        local content = file:read("*a")
        file:close()

        -- Safely decode the JSON back into a Lua table
        local ok, data = pcall(vim.json.decode, content)
        if ok and type(data) == "table" then
            return data
        end
    end

    -- If the file doesn't exist yet, or the JSON is broken, return an empty array
    return {}
end

return M
