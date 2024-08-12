import { Routes, Route } from 'react-router-dom'
import { Home } from './Home'
import { SignUp } from './SignUp'

const Main = () => (
    <Routes>
        <Route
            path='/home'
            element={
                <>
                    <Home />
                </>
            }
        ></Route>
        <Route
            path='/'
            element={
                <>
                    <SignUp />
                </>
            }
        ></Route>
    </Routes>
)

export default Main
