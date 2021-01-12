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
  theme: {},
  variants: {},
  plugins: [],
};
