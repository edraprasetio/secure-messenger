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

export const SignIn = () => {
    const [username, setUsername] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const [error, setError] = useState<string | null>(null)
    const [errorMessage, setErrorMessage] = useState('')

    const login = async (e: SyntheticEvent) => {
        setErrorMessage('')
        e.preventDefault()

        const url = 'http://localhost:8080' + '/login'

        await fetch(url, {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        })
            .then((response) => {
                console.log(response)
                if (response.ok) {
                    return response.json()
                } else {
                    throw new Error('Invalid username or password')
                }
            })
            .then((data) => {
                if (!data.token) {
                    throw new Error('Invalid username or password')
                }
                localStorage.setItem('jwt', data.token)
                localStorage.setItem('username', String(username))
                console.log(localStorage)
            })
            .catch((error) => {
                console.log(error)
                if (error.message) {
                    setErrorMessage(error.message)
                } else {
                    setErrorMessage('An error has occurred')
                }
            })
    }

    // const handleSubmit = (e: React.FormEvent) => {
    //     e.preventDefault()

    //     setError(null)
    //     console.log('Form Submitted:', { username, password })
    //     // Here you can add your logic to send data to your backend
    // }

    return (
        <div>
            <HomeBackground>
                <H1>Log into the Secure Messenger</H1>
                <form onSubmit={login}>
                    <Input type='username' placeholder='Username' value={username} onChange={(e) => setUsername(e.target.value)} required />
                    <Input type='password' placeholder='Password' value={password} onChange={(e) => setPassword(e.target.value)} required />
                    <BlueButton type='submit'>
                        <H1>Sign In</H1>
                    </BlueButton>
                </form>
            </HomeBackground>
        </div>
    )
}
