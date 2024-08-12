import styled from '@emotion/styled'

const BaseButton = styled.button`
    padding: 0px 24px;
    border-radius: 8px;
    align-items: center;
    justify-content: center;
    line-height: 29.26px;
    transition: border-color 0.3s ease;
`

export const ClearButton = styled(BaseButton)`
    border: 2px solid rgba(0, 0, 0, 0);
    background-color: transparent;
    color: ${(props) => props.theme.primaryColor.black[1]};
    &:hover {
        border-color: ${(props) => props.theme.primaryColor.black[1]};
    }
    &:active {
        color: ${(props) => props.theme.primaryColor.white[1]};
        background-color: ${(props) => props.theme.primaryColor.black[1]};
    }
`

export const BlueButton = styled(BaseButton)`
    width: 240px;
    padding: 4px 24px;
    border: 2px solid ${(props) => props.theme.primaryColor.blue[1]};
    background-color: ${(props) => props.theme.primaryColor.blue[1]};
    color: ${(props) => props.theme.primaryColor.white[1]};
    &:hover {
        background-color: ${(props) => props.theme.primaryColor.blue[2]};
        border: 2px solid ${(props) => props.theme.primaryColor.blue[2]};
    }
    &:active {
        background-color: ${(props) => props.theme.primaryColor.white[1]};
        border: 2px solid ${(props) => props.theme.primaryColor.blue[1]};
        color: ${(props) => props.theme.primaryColor.blue[1]};
    }
`

export const WhiteButton = styled(BaseButton)`
    width: 240px;
    padding: 4px 24px;
    border: 2px solid ${(props) => props.theme.primaryColor.white[1]};
    background-color: ${(props) => props.theme.primaryColor.white[1]};
    color: ${(props) => props.theme.primaryColor.black[1]};
    &:hover {
        background-color: ${(props) => props.theme.primaryColor.white[3]};
        border: 2px solid ${(props) => props.theme.primaryColor.white[3]};
    }
    &:active {
        background-color: ${(props) => props.theme.primaryColor.black[1]};
        border: 2px solid ${(props) => props.theme.primaryColor.white[1]};
        color: ${(props) => props.theme.primaryColor.white[1]};
    }
`
