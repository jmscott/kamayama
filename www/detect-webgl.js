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

		// Detect WebGL 1
		const canvas1 = document.createElement('canvas');
		const gl1 = canvas1.getContext('webgl') ||
			canvas1.getContext('experimental-webgl');
		const hasWebGL1 = !!gl1;

		// Detect WebGL 2
		const canvas2 = document.createElement('canvas');
		const gl2 = canvas2.getContext('webgl2');
		const hasWebGL2 = !!gl2;

		console.log(`WebGL 1: ${hasWebGL1}, WebGL 2: ${hasWebGL2}`);   
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
