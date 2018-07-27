components {
  id: "tilemap"
  component: "/level/level1_template.tilemap"
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
  data: "collision_shape: \"/level/level1_template.tilemap\"\ntype: COLLISION_OBJECT_TYPE_STATIC\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"obstacle\"\nmask: \"player\"\nmask: \"enemy\"\nmask: \"obstacle\"\nmask: \"projectile\"\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
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
