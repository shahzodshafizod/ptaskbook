rem барои санчидани 28
@echo off
set /p adad="Адади серакама: "
set satr=

set /a a=%adad%/100
set /a b=%adad%/10%%10
set /a c=%adad%%%10

rem садй

if %a%==1 (
	set satr=сад
	if %b%+%c% neq 0 set satr=яксад
) else (
	if %a%==2 (
		set satr=дусад
	) else (
		if %a%==3 (
			set satr=сесад
		) else (
			if %a%==4 (
				set satr=чорсад
			) else (
				if %a%==5 (
					set satr=панчсад
				) else (
					if %a%==6 (
						set satr=шашсад
					) else (
						if %a%==7 (
							set satr=хафтсад
						) else (
							if %a%==8 (
								set satr=хаштсад
							) else (
								if %a%==9 (
									set satr=нухсад
								)
							)
						)
					)
				)
			)
		)
	)
)

rem дахй

if %b% neq 0 (
	if %a% neq 0 set satr=%satr%у 
	if %b%==1 (
		if %c%==0 (
			set satr=%satr%дах
		) else (
			if %c%==1 (
				set satr=%satr%ёздах
			) else (
				if %c%==2 (
					set satr=%satr%дувоздах
				) else (
					if %c%==3 (
						set satr=%satr%сездах
					) else (
						if %c%==4 (
							set satr=%satr%чордах
						) else (
							if %c%==5 (
								set satr=%satr%понздах
							) else (
								if %c%==6 (
									set satr=%satr%шонздах
								) else (
									if %c%==7 (
										set satr=%satr%хабдах
									) else (
										if %c%==8 (
											set satr=%satr%хаждах
										) else (
											if %c%==9 (
												set satr=%satr%нуздах
											)
										)
									)
								)
							)
						)
					)
				)
			)
		)
	) else (
		if %b%==2 (
			set satr=%satr%бист
		) else (
			if %b%==3 (
				set satr=%satr%си
			) else (
				if %b%==4 (
					set satr=%satr%чил
				) else (
					if %b%==5 (
						set satr=%satr%панчох
					) else (
						if %b%==6 (
							set satr=%satr%шаст
						) else (
							if %b%==7 (
								set satr=%satr%хафтод
							) else (
								if %b%==8 (
									set satr=%satr%хаштод
								) else (
									if %b%==9 (
										set satr=%satr%навад
									)
								)
							)
						)
					)
				)
			)
		)
	)
)

rem вохид

if %c% neq 0 (
	if %b% neq 1 (
		set /a temp=%a%+%b%
		if %temp% neq 0 (
			set satr=%satr%у 
		)
		if %c%==1 (
			set satr=%satr%як
		) else (
			if %c%==2 (
				set satr=%satr%ду
			) else (
				if %c%==3 (
					set satr=%satr%се
				) else (
					if %c%==4 (
						set satr=%satr%чор
					) else (
						if %c%==5 (
							set satr=%satr%панч
						) else (
							if %c%==6 (
								set satr=%satr%шаш
							) else (
								if %c%==7 (
									set satr=%satr%хафт
								) else (
									if %c%==8 (
										set satr=%satr%хашт
									) else (
										if %c%==9 (
											set satr=%satr%нух
										)
									)
								)
							)
						)
					)
				)
			)
		)
	)
)

echo %satr%
echo on