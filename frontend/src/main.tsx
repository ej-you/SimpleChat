import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './assets/index.css'
import './assets/reset.css'
import Router from './Router.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Router />
  </StrictMode>,
)
