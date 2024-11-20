import { useCallback, useEffect, useRef, useState } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'
import { useChatStore } from '../../../store/store'

const Footer = () => {
	const addMessage = useChatStore(state => state.addMessage)
	const [value, setValue] = useState('')
	const textareaElement = useRef<HTMLTextAreaElement>(null)
	const { handleSubmit, register, setValue: setFormValue, reset } = useForm<{ content: string }>({
		defaultValues: {
			content: ''
		}
	})

	// Очистка поля, получение данных
	const onSubmit: SubmitHandler<{ content: string }> = useCallback((data, e) => {
		addMessage(data)
		reset()
		e?.target.reset()
	}, [addMessage, reset])

	// Сохранение значений поля
	const handleChange = (e: React.FormEvent<HTMLTextAreaElement>) => {
		const newValue = (e.target as HTMLTextAreaElement).value;
		setValue(newValue);
		setFormValue('content', newValue);
	}

	// Ctrl + shift - перенос строки
	const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
		if (e.key === 'Enter' && !e.shiftKey) {
			e.preventDefault()
		}
	}

	// Сброс высоты текстового поля, чтобы затем установить его на scrollHeight
	useEffect(() => {
		if (textareaElement.current) {
			textareaElement.current.style.height = 'auto'
			textareaElement.current.style.height = `${textareaElement.current.scrollHeight}px`
		}
	}, [value, onSubmit])
	
	return (
		<footer className='flex flex-col gap-4 background-400'>
			<hr className='w-full border-background-400' />
			<form className='flex gap-4' onSubmit={handleSubmit(onSubmit)}>
				<textarea
					{...register('content')}
					ref={textareaElement}
					tabIndex={0}
					rows={1}
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

				<button type='submit' className='bg-primary rounded-xl px-3	h-fit aspect-[1/1]'>
					<img className='w-6' src='../../../public/uil_message.svg' alt='' />
				</button>
			</form>
		</footer>
	)
}

export default Footer