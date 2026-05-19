/*
	For more comprehensive detection, including checking for WebGL 2
	support, use getContext('webgl2') and check for the
	WebGL2RenderingContext interface.

	To detect if hardware acceleration is active, retrieve the
	WEBGL_debug_renderer_info extension and check the renderer string for
	software keywords like "swiftshader" or "llvmpipe".

		Basic Check: !!canvas.getContext('webgl')
		WebGL 2 Check: !!canvas.getContext('webgl2')
		Hardware Acceleration: Check
			gl.getParameter(
				gl.UNMASKED_RENDERER_WEBGL)
			for software fallbacks.
		Error Handling:
			Listen for the webglcontextcreationerror event to catch
			cases where context creation fails due to blacklisted
			GPUs or disabled settings.
*/
function detectWebGL() {
    try {
        const canvas = document.createElement('canvas');
        // Check for standard and legacy context names
        const gl = canvas.getContext('webgl') || canvas.getContext('experimental-webgl');
        
        if (gl && typeof gl.getParameter === 'function') {
            return true;
        }
        return false;
    } catch (e) {
        return false;
    }
}
