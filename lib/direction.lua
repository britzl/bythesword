local M = {}

function M.direction(from, to)
	local h = 0
	if from.x > to.x then
		h = -1
	elseif from.x < to.x then
		h = 1
	end
	local v = 0
	if from.y > to.y then
		v = -1
	elseif from.y < to.y then
		v = 1
	end
	return h, v
end

return M