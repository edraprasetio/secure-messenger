import { MediumBlackButton } from '../components/atoms/button'
import { HomeBackground } from '../components/home/background'
import { H1 } from '../styles/typography'

export const SignUp = () => {
    return (
        <div>
            <HomeBackground>
                <H1>Sign Up to the Secure Messenger</H1>
                <MediumBlackButton>Sign Up</MediumBlackButton>
            </HomeBackground>
        </div>
    )
}
