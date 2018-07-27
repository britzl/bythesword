local M = {}

local function distance(p1, p2)
	local pd = p1 - p2
	return math.abs(pd.x * pd.x) + math.abs(pd.y * pd.y)
end

function M.in_range(p1, p2, range)
	local d = distance(p1, p2)
	local r = range * range
	return d < r
end


return M