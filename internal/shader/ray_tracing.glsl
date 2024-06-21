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
const float THRESHOLD = 0.000001;

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

// ----------------------------------------
// Sphere
// ----------------------------------------

struct Sphere {
  vec3 center;
  float radius;
  vec3 color;
};

float SphereClosestDistanceAlongRay(Sphere sphere, Ray ray) {
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

// ----------------------------------------
// Scene
// ----------------------------------------

Sphere spheres[6] = Sphere[6](
  Sphere(vec3(0, 2, 0), 2, vec3(255, 0, 0)),
  Sphere(vec3(7, 0, 2), 2, vec3(255, 0, 255)),
  Sphere(vec3(6, 1, -4), 1, vec3(255, 255, 0)),
  Sphere(vec3(-2, 2, 4), 2, vec3(0, 255, 0)),
  Sphere(vec3(-4, 4, 10), 4, vec3(0, 0, 255)),
  Sphere(vec3(-3.2, 1, -1), 1, vec3(0, 255, 255))
);

struct Scene {
  Camera camera;
};

vec3 trace(Scene scene, float x, float y) {
  Ray ray = rayFor(scene.camera, x, y);

  Sphere nearest = Sphere(vec3(0, 0, 0), 0, vec3(0, 0, 0));
  float shortestDistance = INF;

  for (int i = 0; i < spheres.length(); ++i) {
    float distance = SphereClosestDistanceAlongRay(spheres[i], ray);
    if (distance < shortestDistance) {
      shortestDistance = distance;
      nearest = spheres[i];
    }
  }

  if (nearest.radius == 0) {
    return background;
  }

  return nearest.color;
}

// ----------------------------------------
// Main
// ----------------------------------------

void main() {
  // Camera camera = newCamera(vec3(-5, 7, -15), vec3(0, 0, 0));
  Camera camera = newCamera(vec3(-5, 7, -15), vec3(0, 4, 0));
  Scene scene = Scene(camera);

  ivec2 pos = ivec2(gl_GlobalInvocationID.xy);

  float x = (pos.x / width) - 0.5;
  float y = (pos.y / height) - 0.5;

  imageStore(outTex, pos, vec4(trace(scene, x, y), 0));
}
