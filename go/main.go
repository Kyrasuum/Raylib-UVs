package main

/*
#include "raylib.h"

void UpdateModelUVs(Model* mdl) {
	UpdateMeshBuffer(mdl->meshes[0], 1, &mdl->meshes->texcoords[0], mdl->meshes->vertexCount*2*sizeof(float), 0);
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/gen2brain/raylib-go/raylib"
)

var ()

func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "UV raylib test")

	// Define the camera to look into our 3d world
	camera := rl.Camera{}
	camera.Position = rl.NewVector3(0.0, 1.0, 3.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	obj := rl.LoadModel("square.obj")                            // Load OBJ model
	texture := rl.LoadTexture("square.png")                      // Load model texture
	rl.SetMaterialTexture(obj.Materials, rl.MapDiffuse, texture) // Set map diffuse texture

	position := rl.NewVector3(0.0, 0.0, 0.0) // Set model position

	cycle := float32(0)
	//----------------------------------------------------------------------------------

	// Set render cycle to 60 fps
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		//move UVs
		length := int(obj.Meshes.VertexCount)
		var mdluvs []float32

		header := (*reflect.SliceHeader)(unsafe.Pointer(&mdluvs))
		header.Data = uintptr(unsafe.Pointer(obj.Meshes.Texcoords))
		header.Len = length * 2
		header.Cap = length * 2
		for i := 0; i < length; i++ {
			if mdluvs[i*2] >= 0.5 {
				mdluvs[i*2] = cycle/200.0 + 0.5
			} else {
				mdluvs[i*2] = cycle / 200.0
			}
			if mdluvs[i*2+1] >= 0.5 {
				mdluvs[i*2+1] = cycle/200.0 + 0.5
			} else {
				mdluvs[i*2+1] = cycle / 200.0
			}
		}
		C.UpdateModelUVs((*C.Model)(unsafe.Pointer(&obj)))
		cycle = float32(int(cycle+2) % 100)
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)
		rl.DrawModel(obj, position, 1.0, rl.White) // Draw 3d model with texture
		rl.DrawGrid(20, 10.0)                      // Draw a grid
		rl.EndMode3D()

		rl.DrawText(fmt.Sprintf("%+v", cycle), screenWidth-200, screenHeight-20, 10, rl.Gray)
		rl.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	rl.UnloadTexture(texture) // Unload texture
	rl.UnloadModel(obj)       // Unload model

	rl.CloseWindow()
	//--------------------------------------------------------------------------------------
}
