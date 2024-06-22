#version 460 core

layout(local_size_x = 1, local_size_y = 1, local_size_z = 1) in;
layout(rgba32f, binding = 0) uniform image2D outTex;
layout(std430, binding = 1) buffer size
{
  float width;
  float height;
};

// TODO: add to buffer
const vec3 background = vec3(127, 127, 127);

// ----------------------------------------
// Constants
// ----------------------------------------

const float INF = 1. / 0.;
const float THRESHOLD = 0.0001;

// ----------------------------------------
// Vector
// ----------------------------------------

struct UnitVectors
{
  vec3 O;
  vec3 U;
  vec3 X;
  vec3 Y;
  vec3 Z;
  vec3 XI;
  vec3 YI;
  vec3 ZI;
};

const UnitVectors UNITS = UnitVectors(
  vec3(0, 0, 0),
  vec3(1, 1, 1),
  vec3(1, 0, 0),
  vec3(0, 1, 0),
  vec3(0, 0, 1),
  vec3(-1, 0, 0),
  vec3(0, -1, 0),
  vec3(0, 0, -1)
);

float squid(vec3 vector) {
  return dot(vector, vector);
}

// ----------------------------------------
// Ray
// ----------------------------------------

struct Ray {
  vec3 start;
  vec3 direction;
};

Ray newRay(vec3 start, vec3 direction) {
  return Ray(start, normalize(direction));
}

vec3 reflectRay(Ray ray, vec3 normal) {
  float normalSquid = squid(normal);
  if (normalSquid == 0) {
    return ray.direction;
  }

  return ray.direction - (normal * (2 * dot(ray.direction, normal) / normalSquid));
}

// ----------------------------------------
// Camera
// ----------------------------------------

struct Camera {
  vec3 location;
  vec3 lookAt;
  vec3 direction;
  vec3 right;
  vec3 down;
};

Camera newCamera(vec3 location, vec3 lookAt) {
  vec3 direction = normalize(lookAt - location);
  vec3 right = normalize(cross(UNITS.Y, direction)) * 2.;
  vec3 down = normalize(cross(right, direction)) * 1.125;

  return Camera(location, lookAt, direction, right, down);
}

Ray rayFor(Camera camera, float x, float y) {
  vec3 xRay = camera.right * x;
  vec3 yRay = camera.down * y;
  vec3 rayDirection = camera.direction + xRay + yRay;

  return newRay(camera.location, rayDirection);
}

vec3 pointAtDistance(Ray ray, float distance) {
  return ray.start + (ray.direction * distance);
}

// ----------------------------------------
// Material
// ----------------------------------------

struct Finish {
  float ambient;
  float diffuse;
  float shiny;
};

const Finish defaultFinish = Finish(0.1, 0.7, 0);
const Finish shinyFinish = Finish(0.1, 0.7, 0.5);

vec3 highlightFor(Finish finish, vec3 reflection, vec3 light, vec3 lightColor) {
  if (finish.shiny == 0) {
    return vec3(0, 0, 0);
  }

  float intensity = dot(reflection, normalize(light));
  if (intensity <= 0) {
    return vec3(0, 0, 0);
  }

  float exp = 32 * finish.shiny * finish.shiny;
  intensity = pow(intensity, exp);

  return lightColor * (finish.shiny * intensity);
}

struct Material {
  vec3 _color;
  Finish finish;
};

vec3 ambientColor(Material mat) {
  return mat._color * mat.finish.ambient;
}

vec3 diffuseColor(Material mat) {
  return mat._color * mat.finish.diffuse;
}

// ----------------------------------------
// Sphere
// ----------------------------------------

struct Sphere {
  vec3 center;
  float radius;
  Material material;
};

float sphereClosestDistanceAlongRay(Sphere sphere, Ray ray) {
  vec3 os = ray.start - sphere.center;
  float b = 2. * dot(os, ray.direction);
  float c = squid(os) - sphere.radius*sphere.radius;

  float discriminant = b*b - 4.*c;
  if (discriminant < 0) {
    return INF;
  }
  if (discriminant == 0) {
    return -b / 2.;
  }

  float distance1 = (-b - sqrt(discriminant)) / 2.;
  float distance2 = (-b + sqrt(discriminant)) / 2.;
  if (distance1 > THRESHOLD && distance1 < distance2) {
    return distance1;
  }
  if (distance2 > THRESHOLD) {
    return distance2;
  }

  return INF;
}

vec3 sphereNormalAt(Sphere sphere, vec3 point) {
  return normalize(point + (-1 * sphere.center));
}

// ----------------------------------------
// Plane
// ----------------------------------------

struct Plane {
  vec3 normal;
  float distance;
  Material material;
};

float planeClosestDistanceAlongRay(Plane plane, Ray ray) {
  float a = dot(ray.direction, plane.normal);
  if (a == 0) {
    return INF;
  }

  float b = dot(plane.normal, ray.start + (plane.normal * -plane.distance));
  float distance = -b / a;
  if (distance > THRESHOLD) {
    return distance;
  }

  return INF;
}

// ----------------------------------------
// Box
// ----------------------------------------

struct Box {
  vec3 lowerCorner;
  vec3 upperCorner;
  Material material;
};

float boxClosestDistanceAlongRay(Box box, Ray ray) {
  float distances[6] = float[6](INF, INF, INF, INF, INF, INF);

  vec3 lower = (box.lowerCorner - ray.start) / ray.direction;
  vec3 upper = (box.upperCorner - ray.start) / ray.direction;

  if (ray.direction.x != 0) {
    vec3 point = ray.start + (ray.direction * lower.x);
    if (box.lowerCorner.y < point.y && point.y < box.upperCorner.y
      && box.lowerCorner.z < point.z && point.z < box.upperCorner.z
    ) {
      distances[0] = lower.x;
    }

    point = ray.start + (ray.direction * upper.x);
    if (box.lowerCorner.y < point.y && point.y < box.upperCorner.y
      && box.lowerCorner.z < point.z && point.z < box.upperCorner.z
    ) {
      distances[1] = upper.x;
    }
  }

  if (ray.direction.y != 0) {
    vec3 point = ray.start + (ray.direction * lower.y);
    if (box.lowerCorner.x < point.x && point.x < box.upperCorner.x
      && box.lowerCorner.z < point.z && point.z < box.upperCorner.z
    ) {
      distances[2] = lower.y;
    }

    point = ray.start + (ray.direction * upper.y);
    if (box.lowerCorner.x < point.x && point.x < box.upperCorner.x
      && box.lowerCorner.z < point.z && point.z < box.upperCorner.z
    ) {
      distances[3] = upper.y;
    }
  }

  if (ray.direction.z != 0) {
    vec3 point = ray.start + (ray.direction * lower.z);
    if (box.lowerCorner.y < point.y && point.y < box.upperCorner.y
      && box.lowerCorner.x < point.x && point.x < box.upperCorner.x
    ) {
      distances[4] = lower.z;
    }

    point = ray.start + (ray.direction * upper.z);
    if (box.lowerCorner.y < point.y && point.y < box.upperCorner.y
      && box.lowerCorner.x < point.x && point.x < box.upperCorner.x
    ) {
      distances[5] = upper.z;
    }
  }

  float shortest = INF;
  for (int i = 0; i < 6; ++i) {
    if (distances[i] < shortest && distances[i] > THRESHOLD) {
      shortest = distances[i];
    }
  }

  return shortest;
}

vec3 boxNormalAt(Box box, vec3 point) {
  vec3 lowerDiff = abs(box.lowerCorner - point);
  if (lowerDiff.x < THRESHOLD) {
    return UNITS.XI;
  }
  if (lowerDiff.y < THRESHOLD) {
    return UNITS.YI;
  }
  if (lowerDiff.z < THRESHOLD) {
    return UNITS.ZI;
  }

  vec3 upperDiff = abs(box.upperCorner - point);
  if (upperDiff.x < THRESHOLD) {
    return UNITS.X;
  }
  if (upperDiff.y < THRESHOLD) {
    return UNITS.Y;
  }
  if (upperDiff.z < THRESHOLD) {
    return UNITS.Z;
  }

  return UNITS.O;
}

// ----------------------------------------
// Light
// ----------------------------------------

struct Light {
  vec3 position;
  vec3 color;
};

// ----------------------------------------
// Scene
// ----------------------------------------

Sphere spheres[5] = Sphere[5](
  Sphere(vec3(7, 0, 2), 2, Material(vec3(255, 0, 255), shinyFinish)),
  Sphere(vec3(6, 1, -4), 1, Material(vec3(255, 255, 0), shinyFinish)),
  Sphere(vec3(-2, 2, 4), 2, Material(vec3(0, 255, 0), shinyFinish)),
  Sphere(vec3(-4, 4, 10), 4, Material(vec3(0, 0, 255), shinyFinish)),
  Sphere(vec3(-3.2, 1, -1), 1, Material(vec3(0, 255, 255), shinyFinish))
);

Plane planes[1] = Plane[1](Plane(UNITS.Y, 0, Material(vec3(255, 255, 255), defaultFinish)));

Box boxes[1] = Box[1](Box(vec3(-2, 0, -2), vec3(2, 4, 2), Material(vec3(255, 0, 0), shinyFinish)));

Light lights[1] = Light[1](Light(vec3(-30, 25, -12), vec3(255, 255, 255)));

struct Scene {
  Camera camera;
};

bool inShadow(vec3 point, vec3 light) {
  Ray ray = newRay(point, light);
  float lenght = length(light);

  for (int i = 0; i < spheres.length(); ++i) {
    if (sphereClosestDistanceAlongRay(spheres[i], ray) <= lenght) {
      return true;
    }
  }

  for (int i = 0; i < boxes.length(); ++i) {
    if (boxClosestDistanceAlongRay(boxes[i], ray) <= lenght) {
      return true;
    }
  }

  for (int i = 0; i < planes.length(); ++i) {
    if (planeClosestDistanceAlongRay(planes[i], ray) <= lenght) {
      return true;
    }
  }

  return false;
}

vec3 colorAt(Scene scene, vec3 point, Material material, vec3 normal, Ray ray) {
  vec3 color = ambientColor(material);
  vec3 reflectionVec = reflectRay(ray, normal);
  for (int i = 0; i < lights.length(); ++i) {
    vec3 lightVector = lights[i].position - point;
    if (inShadow(point, lightVector)) {
      continue;
    }

    float brightness = dot(normal, normalize(lightVector));
    if (brightness <= 0) {
      continue;
    }

    vec3 illumination = clamp(diffuseColor(material) * lights[i].color, 0, 255) * brightness;
    color = clamp(color + illumination, 0, 255);

    vec3 highlight = highlightFor(material.finish, reflectionVec, lightVector, lights[i].color);
    color = clamp(color + highlight, 0, 255);
  }

  return color;
}

vec3 trace(Scene scene, float x, float y) {
  Ray ray = rayFor(scene.camera, x, y);

  int nearestIdx = -1;
  int nearestType = -1;
  float shortestDistance = INF;

  for (int i = 0; i < spheres.length(); ++i) {
    float distance = sphereClosestDistanceAlongRay(spheres[i], ray);
    if (distance < shortestDistance) {
      shortestDistance = distance;
      nearestIdx = i;
      nearestType = 0;
    }
  }

  for (int i = 0; i < planes.length(); ++i) {
    float distance = planeClosestDistanceAlongRay(planes[i], ray);
    if (distance < shortestDistance) {
      shortestDistance = distance;
      nearestIdx = i;
      nearestType = 1;
    }
  }

  for (int i = 0; i < boxes.length(); ++i) {
    float distance = boxClosestDistanceAlongRay(boxes[i], ray);
    if (distance < shortestDistance) {
      shortestDistance = distance;
      nearestIdx = i;
      nearestType = 2;
    }
  }

  if (nearestIdx == -1) {
    return background;
  }

  vec3 point = pointAtDistance(ray, shortestDistance);
  vec3 normalAt;
  Material material;

  if (nearestType == 0) {
    material = spheres[nearestIdx].material;
    normalAt = sphereNormalAt(spheres[nearestIdx], point);
  } else if (nearestType == 1) {
    material = planes[nearestIdx].material;
    normalAt = planes[nearestIdx].normal;
  } else {
    material = boxes[nearestIdx].material;
    normalAt = boxNormalAt(boxes[nearestIdx], point);
  }

  return colorAt(scene, point, material, normalAt, ray);
}

// ----------------------------------------
// Main
// ----------------------------------------

void main() {
  Camera camera = newCamera(vec3(-5, 7, -15), vec3(0, 4, 0));
  Scene scene = Scene(camera);

  ivec2 pos = ivec2(gl_GlobalInvocationID.xy);

  float x = (pos.x / width) - 0.5;
  float y = (pos.y / height) - 0.5;

  imageStore(outTex, pos, vec4(trace(scene, x, y), 0));
}
