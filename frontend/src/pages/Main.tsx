import { Routes, Route } from 'react-router-dom'
import { Home } from './Home'
import { SignUp } from './SignUp'
import { SignIn } from './SignIn'

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
            path='/login'
            element={
                <>
                    <SignIn />
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
