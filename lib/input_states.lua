local M = {}

M.actions = {}


function M.is_pressed(action_id)
	return M.actions[tostring(action_id)] == true
end

function M.on_input(action_id, action)
	if action.pressed then
		M.actions[tostring(action_id)] = true
	elseif action.released then
		M.actions[tostring(action_id)] = false
	end
end


return M