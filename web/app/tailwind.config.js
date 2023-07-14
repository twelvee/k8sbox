module.exports = {
    content: ['./src/**/*.{vue,js,ts}'],
    plugins: [require('daisyui')],
    daisyui: {
        themes: ["night"],
    },
    theme: {
        colors: {
            transparent: 'transparent',
            current: 'currentColor',
            'white': '#ffffff',
            'purple': '#3f3cbb',
            'midnight': '#0F1623',
            'midnight-light': '#111827',
            'metal': '#565584',
            'tahiti': '#3ab7bf',
            'silver': '#ecebff',
            'moonlight': '#9CA3AF',
            'bubble-gum': '#ff77e9',
            'bermuda': '#78dcca',
        },
    },
};
