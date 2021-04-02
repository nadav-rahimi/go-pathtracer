# Notes
These are notes made during the creation of the project, all code adapted from the book *Ray Tracing in One Weekend* by *Peter Shirley*

### General
- The core of the tracer is to send rays through pixels and compute what colour is seen in the direction of those rays, i.e.
    - calculate which ray goes from the eye to the pixel
    - compute what the ray intersects
    - compute the colour for that intersection point

### The Camera
- The camera centre (eye) is put at `(0, 0, 0)`
- Y-axis goes up and X-axis goes to the right
- Into the screen is the negative z-axis (in order to respect a right-handed coordinate system)
- The screen is traversed from the lower left hand corner and two offset vectors are used to move the ray endpoint across the screen

### A Sphere
- Sphere centre `C` denoted as `cx, cy, cz`
- Point `P` denoted as`x, y, z` is on the sphere if `(x-cx)*(x-cx) + (y-cy)*(y-cy) + (z-cz)*(z-cz) = R*R`
    - The dot product form of this is `dot((P-C), (P-C)) = R*R`
    - When `P(t)` becomes the function of where the ray is at a given time, the equation becomes `dot((A + t*B - C), (A + t*B - C)) = R*R`
        - Expanding this into a general equation we get `t*t*dot(B,B) + 2*t*dot(A-C,A-C) + dot(C,C) - R*R = 0`
        - Solving this quadratic equation, zero roots means no collision with sphere, 1 root means collision at tangent, 2 roots means collision at two different points on the sphere

### Surface Normals & Multiple Objects
- Surface normals are used to shade
    - These are vectors which are perpendicular to the surface
    - By convention these vectors point out
    - One design decision is to make these normals unit vectors, this is convenient for shading so it is recommended
- For spheres the normal is the direction of the hitpoint minus the centre
- Normals are visualised with a colour map, since they are -1 to 1 each component can be mapped to an rgb value if he restrict it to the interval 0 to 1. Also to shade the closes hitpoint, not the further one which is behind it we need the smallest "t".

### Anti-Aliasing
- Aliasing averages samples inside each pixel to reduce jagged edges
- Random value generated in the range 0 <= val < 1

### Diffuse Materials
- Do not emit light, they take on the colour of their surroundings, they modulate it with their own intrinsic colour
- Light that reflects off a diffuse surface has a randomised direction
- Some rays may be absorbed rather than refelcted, the darker the surface the more likely this happens
- Pick a random point **S** from the unit radius sphere that is tangent to the hitpoint and send a ray from the hitpoint **p** to the random point **s**. Sphere has a centre (**p** + **N**)

### Metals 
- These reflect, only render them if the reflection exists or stack overflow occurs
- Add fuzziness to make them more realistic

### Dialectrics
- When light rays hit them they split into refracted and reflected rays
- Refract using snell's law, but make sure to take into account for total internal reflection

### Camera FOV
- This implementation specifies it in degrees and changes it into radians
- Standard FOV is 90 degrees
- Camera looks from a point, to a point