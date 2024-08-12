import { BlueButton } from '../components/atoms/button'
import { HomeBackground } from '../components/home/background'
import { H1 } from '../styles/typography'

export const SignUp = () => {
    return (
        <div>
            <HomeBackground>
                <H1>Sign Up to the Secure Messenger</H1>
                <BlueButton>
                    <H1>Sign Up</H1>
                </BlueButton>
            </HomeBackground>
        </div>
    )
}
