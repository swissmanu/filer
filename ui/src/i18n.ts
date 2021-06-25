import { addMessages, getLocaleFromNavigator, init } from "svelte-i18n";
import de from "./locales/de.json";
import en from "./locales/en.json";

addMessages("en", en);
addMessages("de", de);

init({
  fallbackLocale: "en",
  initialLocale: getLocaleFromNavigator(),
});
