#version 460 core

layout(local_size_x = 1, local_size_y = 1, local_size_z = 1) in;
layout(rgba32f, binding = 0) uniform image2D outTex;
layout(std430, binding = 1) buffer size
{
    float width;
    float height;
};

void main()
{
    ivec2 pos = ivec2(gl_GlobalInvocationID.xy);
    imageStore(outTex, pos, vec4(pos.x/width*255, 0, pos.y/height*255, 0));
}
