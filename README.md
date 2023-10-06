# p-eye-ndi
 use normal cameras to track and zoom given subjects 

**Draft1**
crop the input to a max of 1.5X
detect faces, pan the matrix to keep an average center
auto croup out as needed to keep the faces 

send as an NDI display out
 
 Creating an NDI (Network Device Interface) tool that reads camera feeds and outputs them as NDI streams is a great idea for enhancing the flexibility of your video production workflow. NDI is a popular protocol for sending high-quality video and audio over IP networks, making it easier to share video sources across different software and hardware. Here's a step-by-step guide on how to create such a tool:

**Prerequisites:**

1. **Development Environment**: You will need a programming environment suitable for your chosen programming language. NDI SDK supports various languages like C++, C#, and Python.

2. **NDI SDK**: Download and install the NewTek NDI SDK from the official NDI website (https://www.ndi.tv/sdk/). The SDK includes libraries and documentation that you'll need to work with NDI.

3. **Camera Hardware**: You'll need one or more cameras that you want to capture video from.

**Steps to Create Your NDI Tool:**

1. **Set Up Your Development Environment:**
   
   Choose a programming language for your NDI tool. The NDI SDK supports multiple languages, but C# and C++ are commonly used. Ensure that you have a working development environment for your chosen language.

2. **Install and Configure NDI SDK:**

   - Unzip the NDI SDK files to a directory on your development machine.
   - Follow the installation instructions provided in the SDK documentation for your specific language. This usually involves adding the SDK libraries to your project and configuring build settings.

3. **Capture Video from Cameras:**

   Use the appropriate libraries or APIs to capture video from your cameras. The method for doing this depends on the programming language you're using and the type of cameras you have. Libraries like OpenCV are commonly used for video capture in C++ and Python.

4. **Initialize and Send NDI Streams:**

   - Initialize NDI using the SDK functions for your chosen programming language.
   - Create NDI video frames from the captured video frames.
   - Send these NDI video frames as NDI sources.

   The NDI SDK documentation provides sample code and functions for creating NDI sources and sending video and audio data.

5. **Handle Error Cases:**

   Implement error handling to gracefully handle situations where cameras are unavailable or NDI initialization fails.

6. **Testing:**

   Test your NDI tool on a network. Ensure that it's sending NDI streams that other software can detect and use.

7. **Documentation:**

   Create documentation for your NDI tool, including usage instructions and any dependencies.

8. **Distribution:**

   If you intend to share your tool with others, package it appropriately and distribute it. Consider open-sourcing it if you want to encourage community contributions.

9. **Maintenance and Updates:**

   Keep your NDI tool up-to-date, addressing any bug fixes or adding new features as needed.

Remember that the specifics of implementation may vary depending on the programming language you choose and the exact requirements of your project. Consult the NDI SDK documentation for detailed information on using NDI in your chosen language. Additionally, consider any licensing requirements associated with the NDI SDK and ensure that you comply with them in your project.