components {
  id: "ai_idle"
  component: "/ecs/idle.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "IDLE_DISTANCE"
    value: "0.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "IDLE_ANIMATION"
    value: "orcwiz_idle"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "ai_orcwiz"
  component: "/ecs/ai_orcwiz.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "CAST_RANGE"
    value: "190.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "CAST_ANIMATION"
    value: "orcwiz_attack"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "CAST_EFFECT_FACTORY"
    value: "level:/factories#poisonspikefactory"
    type: PROPERTY_TYPE_URL
  }
  properties {
    id: "CAST_EFFECT_DELAY"
    value: "0.2"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "animation"
  component: "/ecs/animation.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "damage"
  component: "/ecs/hit_damage.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "death_overlay"
  component: "/ecs/death_overlay.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "FACTORY_URL"
    value: "level:/factories#bloodfactory"
    type: PROPERTY_TYPE_URL
  }
}
components {
  id: "health"
  component: "/ecs/health.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "HEALTH"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "CORPSE_FACTORY"
    value: "level:/factories#corpsefactory"
    type: PROPERTY_TYPE_URL
  }
  properties {
    id: "CORPSE_ANIMATION"
    value: "orcwiz_corpse"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "HIT_ANIMATION"
    value: "orcwiz_hit"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "movement"
  component: "/ecs/movement.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "MAX_RUNNING_SPEED"
    value: "430.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "TURN_ACCELERATION"
    value: "15000.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "WALK_ANIMATION"
    value: "orcwiz_run"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "RUN_ANIMATION"
    value: "orcwiz_run"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "script"
  component: "/enemy/enemy.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "sprite_flip"
  component: "/ecs/sprite_flip.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"obstacle\"\n"
  "mask: \"player\"\n"
  "mask: \"weapon\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 7.0\n"
  "      y: 16.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 17.0\n"
  "  data: 14.5\n"
  "  data: 10.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/atlas/game.atlas\"\n"
  "default_animation: \"orcwiz_idle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 25.5
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
