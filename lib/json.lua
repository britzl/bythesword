-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

json_lib = {}

local function is_array(x)
	local is_positive_int = function (n)
		return type(n) == "number" and n > 0 and math.floor(n) == n
	end
	local max, n = 0, 0
	for k,_ in pairs(x) do
		if not is_positive_int(k) then return false end
		max = math.max(max, k)
		n = n + 1
	end
	return n == max
end


local function encode(x) 
	local t = type(x)
	if t == "string" then
		return '"' .. x .. '"'
	elseif t == "boolean" then
		if x then
			return "true"
		else
			return "false"
		end
	elseif t == "table" then
		if is_array(x) then
			local ret = "["
			for i,v in ipairs(x) do
				if i > 1 then
					ret = ret .. ","
				end
				ret = ret .. json_encode(v)
			end
			return ret .. "]"
		else
			local ret = "{"
			local i = 1
			for k,v in pairs(x) do
				if i > 1 then
					ret = ret .. "," 
				end
				ret = ret .. '"' .. k .. '": ' .. json_encode(v)
				i = i + 1  
			end
			return ret .. "}"
		end
	else
		return x
	end
end


json_lib.encode = encode