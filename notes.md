# Current tasks
---

[] Bugs:
    [x] Jumping off of non solid platforms and Falling through platforms while 
       while continually holding the jump button triggers a jump.
       - expected behavior is to land on the platform and not perform a jump.
       - Jump buffer is being set from the initial jump instead from
         subsequent presses during the jump. 
       - Happens when jumping from non solid platforms.
    [x] fix jump animation not working

[] Add coyote time
    [] add can jump variable
    [] Bug: jumping sets coyote timer

    if not grounded and not jumping, set coyote timer
    if coyote timer greater than 0, falling animation
    if coyote timer greater than 0 and jump is pressed, perform jump
    each loop, decrement coyote timer

[x] Add tilemap

[x] Improve player movement
    [x] add jump buffer
    [x] Allow player to control jump height

[x] Animation State machine
    [x] Update Animation component to include the enitity's state 

[x] Camera
    [x] implement camera component
    [x] camera feature
    [x] add camera features to scene
    [x] add system so that camera follows player
    [x] add camera zooming
    [x] make the player the target of the zoom mechanic
    [x] create camera trap
    [x] Camera teleports when player switches direction. Add linear interpolation
        from current camera position to derised (based off the camera trap) position

[x] Add animations
    [x] load texture
    [x] create animation frames from texture
    [x] implement system to advance every animation
    [x] implement draw animation 
    [x] add animation to player 

[x] central asset loading
    [x] create package 
    [x] move font loading to assets


# Temp
---
Update player algorithm:

Add gravity to speedY

if wall sliding and we are falling
    set speedY to max speed

if NOT wall sliding
    apply x movement

apply friction if speed is greater than friction

clamp x speed

if jumping
    if keydown is pressed and on ground and the ground reference is a platform
        set player.IgnorePlatform to the ground reference
    else 
        // perform jump
        apply jumpspd to speedY


---Notes



---Questions
1. When does the WallSliding variable gets set
    
## Changes
- Added animations
- Replaced shapes with assets
- Added cameras
- Added camera system which provides camera functionalities such as zooming
- Added camera follow system which follows the player 
