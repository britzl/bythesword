local M = {}

M.max = 40

local gameobjects = {}

function M.add(go_url)
	table.insert(gameobjects, go_url)
	if #gameobjects > M.max then
		local url = table.remove(gameobjects, 1)
		go.delete(url)
	end
end

return M