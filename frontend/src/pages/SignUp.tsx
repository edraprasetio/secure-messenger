import { useState } from 'react'
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
    const [error, setError] = useState<string | null>(null)

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault()

        if (password !== confirmPassword) {
            setError('Passwords do not match')
            return
        }

        setError(null)
        console.log('Form Submitted:', { email, password })
        // Here you can add your logic to send data to your backend
    }

    return (
        <div>
            <HomeBackground>
                <H1>Sign Up to the Secure Messenger</H1>
                <form onSubmit={handleSubmit}>
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
