import styled from '@emotion/styled'

export const H1 = styled.div`
    font-family: PlusJakartaSans-SemiBold;
    font-size: 16px;
    font-weight: 500;
    letter-spacing: 1px;
`

export const H2 = styled.div`
    font-family: PlusJakartaSans-SemiBold;
    font-size: 20px;
    font-weight: 500;
    letter-spacing: 3px;
`

export const H3 = styled.div`
    font-family: PlusJakartaSans-SemiBold;
    font-size: 32px;
    font-weight: 500;
    letter-spacing: 1px;

    @media (max-width: ${(props) => props.theme.breakPoints.phone}) {
        font-size: 24px;
    }
`

export const H4 = styled.div`
    font-family: PlusJakartaSans-Regular;
    font-size: 16px;
    font-weight: 500;
    letter-spacing: 1px;
    line-height: 1.6;
`

export const H5 = styled.div`
    font-family: PlusJakartaSans-Bold;
    font-size: 20px;
    font-weight: 500;
    letter-spacing: 1px;
    line-height: 1.6;
`
