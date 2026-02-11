components {
  id: "enemy"
  component: "/game/entities/enemy.script"
  properties {
    id: "target"
    value: "/player/player"
    type: PROPERTY_TYPE_URL
  }
}
components {
  id: "health_bar"
  component: "/game/gui/health_bar.gui"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"enemy\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprites/spritesheet_tilesource.tilesource\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 15.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "damage_sound"
  type: "sound"
  data: "sound: \"/assets/audio/EnemyDamaged.wav\"\n"
  ""
}
