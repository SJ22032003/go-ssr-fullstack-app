const { build } = require("esbuild");
const path = require("path");
const glob = require("glob");

async function bundle() {
  try {
    const files = glob.sync("./views/**/*.{js,jsx,ts,tsx}"); // Match all JavaScript, TypeScript, JSX, and TSX files recursively in the 'views' directory
    await build({
      // entryPoints: files, // Use the matched files as entry points
      stdin: { contents: "" }, // Use an empty string as the input file
      inject: files, // Inject all matched files
      bundle: true,
      outfile: path.resolve(__dirname, "static/js/bundle.js"), // Output file
      minify: true, // Optionally minify the output
    });
    console.log("Build completed successfully!");
  } catch (err) {
    console.error("Build failed:", err);
    process.exit(1);
  }
}

bundle();

// const esbuild = require('esbuild');
// const glob = require('glob');

// esbuild
//     .build({
//         stdin: { contents: '' },
//         inject: glob.sync('src/js/**/*.js'),
//         bundle: true,
//         sourcemap: true,
//         minify: true,
//         outfile: 'web/js/bundle.js',
//     })
//     .then(() => console.log("⚡ Javascript build complete! ⚡"))
//     .catch(() => process.exit(1))