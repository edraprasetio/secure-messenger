import { SyntheticEvent, useState } from 'react'
import { BlueButton } from '../components/atoms/button'
import { HomeBackground } from '../components/home/background'
import { H1 } from '../styles/typography'
import styled from '@emotion/styled'

const Input = styled.input`
    width: 100%;
    padding: 10px;
    margin: 8px 0;
    border-radius: 4px;
    border: 1px solid #ccc;
    font-size: 16px;
`

export const SignUp = () => {
    const [username, setUsername] = useState<string>('')
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const [confirmPassword, setConfirmPassword] = useState<string>('')
    const [errorMessage, setErrorMessage] = useState('')

    const register = async (e: SyntheticEvent) => {
        setErrorMessage('')
        e.preventDefault()

        const url = 'http://localhost:8080' + '/register'

        await fetch(url, {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                email: email,
                password: password,
            }),
        }).then((response) => {
            console.log(response)
            if (response.ok) {
                return response.json()
            } else {
                throw new Error('Invalid username or password')
            }
        })
    }

    return (
        <div>
            <HomeBackground>
                <H1>Sign Up to the Secure Messenger</H1>
                <form onSubmit={register}>
                    <Input type='username' placeholder='Username' value={username} onChange={(e) => setUsername(e.target.value)} required />
                    <Input type='email' placeholder='Email' value={email} onChange={(e) => setEmail(e.target.value)} required />
                    <Input type='password' placeholder='Password' value={password} onChange={(e) => setPassword(e.target.value)} required />
                    <Input type='password' placeholder='Confirm Password' value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} required />
                    <BlueButton type='submit'>
                        <H1>Sign Up</H1>
                    </BlueButton>
                </form>
            </HomeBackground>
        </div>
    )
}
