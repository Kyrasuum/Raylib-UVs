#include "raylib.h"
#include <stdio.h>

//------------------------------------------------------------------------------------
// Program main entry point
//------------------------------------------------------------------------------------
int main(void) {
    // Initialization
    //--------------------------------------------------------------------------------------
    const int screenWidth = 800;
    const int screenHeight = 450;

    InitWindow(screenWidth, screenHeight, "UV test");

    // Define the camera to look into our 3d world
    Camera camera = { 0 };
    camera.position = (Vector3){ 0.0f, 1.5f, 5.0f };
    camera.target = (Vector3){ 0.0f, 0.0f, 0.0f };
    camera.up = (Vector3){ 0.0f, 1.0f, 0.0f };
    camera.fovy = 45.0f;
    camera.projection = CAMERA_PERSPECTIVE;

    // Load the object model
    Model model = LoadModel("square.obj");
    Texture2D texture = LoadTexture("square.png");
    model.materials[0].maps[MATERIAL_MAP_DIFFUSE].texture = texture;

    int cycle = 0;

    // Set our game to run at 10 frames-per-second
    SetTargetFPS(10);
    //--------------------------------------------------------------------------------------

    // Main game loop
    // Detect window close button or ESC key
    while (!WindowShouldClose()) {
        // Update
        //----------------------------------------------------------------------------------
        UpdateCamera(&camera, CAMERA_FIRST_PERSON);

        //move UVs
    	int length = model.meshes->vertexCount;
    	for (int i = 0; i < length; i++) {
    	    if (model.meshes->texcoords[i*2] >= 0.5) {
        		model.meshes->texcoords[i*2] = cycle/200.0f + 0.5f;
    	    } else {
        		model.meshes->texcoords[i*2] = cycle/200.0f;
        	}
    	    if (model.meshes->texcoords[i*2+1] >= 0.5) {
        		model.meshes->texcoords[i*2+1] = cycle/200.0f + 0.5f;
    	    } else {
        		model.meshes->texcoords[i*2+1] = cycle/200.0f;
    	    }

    	}
    	UpdateMeshBuffer(model.meshes[0], 1, &model.meshes->texcoords[0], length*2*sizeof(float), 0);
    	cycle = (cycle + 2)% 100;
        //----------------------------------------------------------------------------------

        // Draw
        //----------------------------------------------------------------------------------
        BeginDrawing();
            ClearBackground(RAYWHITE);

            BeginMode3D(camera);
                // Draw a grid
                DrawGrid(10, 1.0f);
                // Draw the model
                DrawModel(model, (Vector3){0,0,0}, 1.0f, WHITE);
            EndMode3D();

            DrawText(TextFormat("Cycle: %08i", cycle), 200, 80, 20, RED);
            DrawFPS(10, 10);
        EndDrawing();
        //----------------------------------------------------------------------------------
    }

    // De-Initialization
    //--------------------------------------------------------------------------------------
    UnloadTexture(texture);     // Unload texture
    UnloadModel(model);         // Unload model
    CloseWindow();        // Close window and OpenGL context
    //--------------------------------------------------------------------------------------

    return 0;
}
