/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      './internal/**/*{go,html,js}',
      './static/**/*{go,html,js}',
    ],
    theme: {
        extend: {
            colors: {
                'primary': '#E51636',
                'gray': {
                    100: '#F7F7F7',
                    500: '#999999',
                    900: '#333333',
                }
            },
        },
    },
    plugins: [],
  };
  