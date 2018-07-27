return {
	-- defold generated message
	CONTACT_POINT_RESPONSE = hash("contact_point_response"),
	-- defold generated message. sent when a sprite animation is done
	ANIMATION_DONE = hash("animation_done"),
	-- defold message to change parent for a game object
	SET_PARENT = hash("set_parent"),
	-- defold message sent as a response to a raycast request
	RAY_CAST_RESPONSE = hash("ray_cast_response"),
	
	-- play an animation
	-- @param id (hash)
	-- @param priority (number)
	ANIMATE = hash("animate"),
	
	STOP_ANIMATION = hash("stop_animation"),
	
	-- inflict damage on recipient
	-- @param amount (number)
	DAMAGE = hash("damage"),
	
	-- push back recipient. recipient must have a movement.script
	-- @param normal
	-- @param distance
	PUSHBACK = hash("pushback"),
	
	-- send when requesting to run in a direction
	-- @param direction (-1, 0, 1)
	RUN = hash("run"),
	
	-- send when requesting to walk in a direction
	-- @param direction (-1, 0, 1)
	WALK = hash("walk"),
	
	-- send when requesting to move in a certain direction
	-- @param amount (velocity, vector3)
	MOVE = hash("move"),
	
	STOP_RUNNING = hash("stop_running"),
	
	-- send when requesting to climb up or down
	-- @param direction (-1, 0, 1)
	CLIMB = hash("climb"),
	
	JUMP = hash("jump"),
	ABORT_JUMP = hash("abort_jump"),
	ATTACK = hash("attack"),
	SPECIAL_ATTACK1 = hash("special_attack1"),
	
	-- fire projectile
	-- @param direction (-1, 0, 1)
	FIRE = hash("fire"),
	FIRE_DONE = hash("fire_done"),
	
	-- change facing
	-- @param direction (-1, 0, 1)
	FACE = hash("face"),
	
	-- sent when a target has been hit
	HIT = hash("hit"),
	
	ATTACKING = hash("attacking"),
	ATTACK_DONE = hash("attack_done"),
	WEAPON_ATTACK_DONE = hash("weapon_attack_done"),
	
	-- when dead but not yet removed (death anim and so on is playing)
	DYING = hash("dying"),
	
	-- when dead and about to be removed (after any death anims are played)
	DEAD = hash("dead"),
	
	-- sent to attacker when the target has been killed
	TARGET_DIED = hash("target_died"),
	
	-- sent when the player has died
	PLAYER_DEAD = hash("player_dead"),
	
	-- sent when the health has changed
	-- @param amount
	-- @param health
	HEALTH_CHANGED = hash("health_changed"),
	
	-- sent when the health has decreased
	-- @param amount
	-- @param health
	HEALTH_DECREASED = hash("health_decreased"),
	
	-- sent when the health has increased
	-- @param amount
	-- @param health
	HEALTH_INCREASED = hash("health_decreased"),
	
	RESTART_LEVEL = hash("restart_level"),
	
	-- send when a script should be enabled
	ENABLE_SCRIPT = hash("enable_script"),
	
	-- send when a script should be disabled
	DISABLE_SCRIPT = hash("disable_script"),
	
	-- sent by enemy scripts to the main script to register them (from init function)
	REGISTER = hash("register"),
	
	-- sent by enemy scripts to the main script to unregister them (from final function)
	UNREGISTER = hash("unregister"),
	
	ACTIVE = hash("active"),
	IDLE = hash("idle"),
	
	REQUEST_VELOCITY = hash("ecs_request_velocity"),
	VELOCITY_RESPONSE = hash("ecs_velocity_response"),
}
