import { useEffect, useRef, useState } from 'react'

const TextField = () => {
	const [value, setValue] = useState('')
	const textareaRef = useRef<HTMLTextAreaElement>(null)

	const handleChange = (e: React.FormEvent<HTMLTextAreaElement>) => {
		setValue((e.target as HTMLTextAreaElement).value)
	}

	// Ctrl + shift - перенос строки
	const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
		if (e.key === 'Enter' && !e.shiftKey) {
			e.preventDefault()
		}
	}

	// Сброс высоты текстового поля, чтобы затем установить его на scrollHeight
	useEffect(() => {
		if (textareaRef.current) {
			textareaRef.current.style.height = 'auto'
			textareaRef.current.style.height = `${textareaRef.current.scrollHeight}px`
		}
	}, [value])

	return (
		<textarea
			tabIndex={0}
			rows={1}
			ref={textareaRef}
			style={{ 
				height: 'auto', 
				maxHeight: '200px',
				overflowY: 'auto'
			}}

			className='w-full h-fit resize-none text-white placeholder:text-subtitle-gray font-bold bg-background-400 appearance-none py-3 px-4 rounded-xl border-subtitle-gray outline-none'
			
			placeholder='Type here...'
			onChange={handleChange}
			onKeyDown={handleKeyDown}
			onInput={(e) => {
				const target = e.target as HTMLTextAreaElement
				target.style.height = `${target.scrollHeight}px`
			}}>
		</textarea>
	)
}

export default TextField