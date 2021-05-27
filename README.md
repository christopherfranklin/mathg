# mathg
MathG is a simple math library for 2D and 3D Programming written in Go.

This library is heavily inspired by the C library - [MathC](https://github.com/felselva/mathc). Although it is inspired, the APIs have diverged quite a bit due to differences between C and Go.

**NOTE:** *This repo is under active development. It will be published as v1.0 when the API is locked in.*

## Features

* Vectors - 2D, 3D, 4D
* Quaternions
* Matrices - 2x2, 3x3, 4x4
* Easing Functions 

## Contributions & Development

MathG is open to all contributions. Please open all changes as a pull request. Also feel free to submit any bugs or feedback in the issue tracker.

## Version

MathG is currently under active development an releasing as *alpha*. Once the API is finalize, the first 1.0 version will be released.

## Types

All types in this library are defined as Structs.

All internal values are represented as `float64`

### Vectors

#### 2D
 ```go
 var pos Vec2 = Vec2{0., 0.}
 pos.X == 0
 pos.Y == 0
 ```

 #### 3D
 ```go
 var pos Vec3 = Vec3{0., 0., 0.}
 pos.X == 0
 pos.Y == 0
 pos.Z == 0
 ```

 #### 4D
 ```go
 var pos Vec4 = Vec4{0., 0., 0., 0.}
 pos.X == 0
 pos.Y == 0
 pos.Z == 0
 pos.W == 0
 ```

 ### Quaternions
 A quaternion is a special case of a `Vec4`. It is a representaion of a complex number in 3 Dimensions.

 ```go
 var q Quaternion = Quaternion{0., 0., 0., 0.}
 ```

 ### Matrices

 #### 2x2
 ```go
 var m Mat2 = Mat2{
   0.,0.,
   0.,0.,
 }
 ```

 Accessing a 2x2 matrix is done using the syntax `MXY`.
 ```
 [
   M11, M21,
   M12, M22,
 ]
 ```

 #### 3x3
 ```go
 var m Mat3 = Mat3{
   0.,0.,0.,
   0.,0.,0.,
   0.,0.,0.,
 }
 ```

 Accessing a 3x3 matrix is done using the syntax `MXY`.
 ```
 [
   M11, M21, M31,
   M12, M22, M32,
   M13, M23, M33,
 ]
 ```

 #### 4x4
 ```go
 var m Mat4 = Mat4{
   0.,0.,0.,0.,
   0.,0.,0.,0.,
   0.,0.,0.,0.,
   0.,0.,0.,0.,
 }
 ```

 Accessing a 4x4 matrix is done using the syntax `MXY`.
 ```
 [
   M11, M21, M31, M41,
   M12, M22, M32, M42,
   M13, M23, M33, M43,
   M14, M24, M34, M44,
 ]
 ```

 ## Easing Functions
 Just like the MathC library, I ported over the Easing Functions.

 The easing functions are an implementation of the functions presented in easings.net, useful particularly for animations.

Easing functions take a value inside the range 0.0-1.0 and usually will return a value inside that same range.

## Usage
Creating a Look At matrix:
```go
position := &Vec3{}
target := &Vec3{0., 0., 10.}
up := &Vec3{0., 1., 0.}
var view *Mat4 = position.LookAt(target, up)
```

Creating a Perspective Matrix:
```go
var perspective *Mat4 = Perspective(ToRadians(60), 1., 0.1, 100.)
```

## License

MIT License

Copyright (c) 2021 Christopher Franklin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
