/** @type {import('tailwindcss').Config} */
import colors, {black, transparent, white} from 'tailwindcss/colors'

export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      black,
      transparent,
      white,
      background: {
        '400' : '#3F3F3F',
        '800' : '#313131',
      },
      'primary' : '#C9E956',
      'title' : '#F9FFE9',
      subtitle: {
        'gray': '#7F8178',
        'purple': '#C68DFE',
      }
    },
    fontSize: {
      sm: '0.9rem',
      base: '1.15rem',
      xl: '2.25rem',
    }
  },
  plugins: [],
}

