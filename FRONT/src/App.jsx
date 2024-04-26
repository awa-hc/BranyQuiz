import { Route, Routes } from 'react-router-dom'
import './App.css'
import { Navbar } from './components/Navbar'
import { Login, Home, Friends, Profile, Live } from './pages'
import { ToastContainer } from 'react-toastify'

export const App = () => {

  return (
    <>
      <ToastContainer />
      <Navbar />
      <Routes>
        <Route path='/' element={<Home/>}/>
        <Route path='/friends' element={<Friends/>}/>
        <Route path='/login' element={<Login/>}/>
        <Route path='/profile' element={<Profile/>}/>
        <Route path='/live' element={<Live/>}/>
      </Routes>
    </>
  )
}

