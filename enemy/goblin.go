components {
  id: "ai_follow"
  component: "/ecs/ai_follow.script"
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
    id: "RANGE"
    value: "200.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
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
}
components {
  id: "ai_melee_attack"
  component: "/ecs/ai_melee_attack.script"
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
  properties {
    id: "IDLE"
    value: "goblin_idle"
    type: PROPERTY_TYPE_HASH
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
}
components {
  id: "melee_attack"
  component: "/ecs/melee_attack.script"
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
    id: "GROUP"
    value: "player"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "ANIMATION"
    value: "goblin_attack"
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
    value: "130.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "TURN_ACCELERATION"
    value: "1000.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "pushback"
  component: "/ecs/hit_pushback.script"
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
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_KINEMATIC\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"enemy\"\nmask: \"obstacle\"\nmask: \"player\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_BOX\n    position {\n      x: 0.0\n      y: 11.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 3\n  }\n  data: 11.0\n  data: 11.0\n  data: 10.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
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
  data: "tile_set: \"/atlas/game.atlas\"\ndefault_animation: \"goblin_idle\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 0.0
    y: 12.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}