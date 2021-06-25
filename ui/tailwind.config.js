const production = !process.env.ROLLUP_WATCH;

module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true,
  },
  purge: {
    enabled: production,
    content: ["./src/**/*.svelte", "./public/**/*.html"],
  },
  theme: {
    flex: {
      1: "1 1 0%",
      2: "2 2 0%",
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
