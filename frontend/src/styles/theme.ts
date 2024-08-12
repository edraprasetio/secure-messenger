import { Theme } from '@emotion/react'
/**
 * Implement the theme interface
 */
export const defaultTheme: Theme = {
    primaryColor: {
        beige: {
            1: '#E5D3B3',
        },
        black: {
            1: '#35363D',
        },
        blue: {
            1: '#4D7CF6',
            2: '#6590FF',
        },
        grey: {
            1: '#666464',
        },
        white: {
            1: '#FFFFFF',
            2: '#F3F3F3',
            3: '#DFDFDF',
        },
    },
    breakPoints: {
        tablet: '1420px',
        miniTablet: '1240px',
        largePhone: '940px',
        phone: '700px',
    },
}
