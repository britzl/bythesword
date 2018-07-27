require "lib.json"

level = {}


local function load(tilemap_url)
	local level_as_string = sys.load_resource("/level/test.json")
	if not level_as_string then return end
	local l = json.decode(level_as_string)
	
	local width = l.width		
	local height = l.height		
	
	local x, y, w, h = tilemap.get_bounds("#tilemap")
	print(x, y, w, h)
	for _,layer in pairs(l.layers) do
		local data = layer.data
		local name = layer.name
		local name_hash = hash(name)
		
		for ix = 1, w do
			for iy = 0, h - 1 do
				local i = ix + (iy * width)
				local tile = data[i]
				local tx = ix
				local ty = h - iy
				tilemap.set_tile(tilemap_url, name, tx, ty, tile)
			end
		end
	end
end

level.load = load