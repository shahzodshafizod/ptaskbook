rem ����� ��������� 28
@echo off
set /p adad="����� ��������: "
set satr=

set /a a=%adad%/100
set /a b=%adad%/10%%10
set /a c=%adad%%%10

rem ����

if %a%==1 (
	set satr=���
	if %b%+%c% neq 0 set satr=�����
) else (
	if %a%==2 (
		set satr=�����
	) else (
		if %a%==3 (
			set satr=�����
		) else (
			if %a%==4 (
				set satr=������
			) else (
				if %a%==5 (
					set satr=�������
				) else (
					if %a%==6 (
						set satr=������
					) else (
						if %a%==7 (
							set satr=�������
						) else (
							if %a%==8 (
								set satr=�������
							) else (
								if %a%==9 (
									set satr=������
								)
							)
						)
					)
				)
			)
		)
	)
)

rem ����

if %b% neq 0 (
	if %a% neq 0 set satr=%satr%� 
	if %b%==1 (
		if %c%==0 (
			set satr=%satr%���
		) else (
			if %c%==1 (
				set satr=%satr%�����
			) else (
				if %c%==2 (
					set satr=%satr%��������
				) else (
					if %c%==3 (
						set satr=%satr%������
					) else (
						if %c%==4 (
							set satr=%satr%������
						) else (
							if %c%==5 (
								set satr=%satr%�������
							) else (
								if %c%==6 (
									set satr=%satr%�������
								) else (
									if %c%==7 (
										set satr=%satr%������
									) else (
										if %c%==8 (
											set satr=%satr%������
										) else (
											if %c%==9 (
												set satr=%satr%������
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
			set satr=%satr%����
		) else (
			if %b%==3 (
				set satr=%satr%��
			) else (
				if %b%==4 (
					set satr=%satr%���
				) else (
					if %b%==5 (
						set satr=%satr%������
					) else (
						if %b%==6 (
							set satr=%satr%����
						) else (
							if %b%==7 (
								set satr=%satr%������
							) else (
								if %b%==8 (
									set satr=%satr%������
								) else (
									if %b%==9 (
										set satr=%satr%�����
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

rem �����

if %c% neq 0 (
	if %b% neq 1 (
		set /a temp=%a%+%b%
		if %temp% neq 0 (
			set satr=%satr%� 
		)
		if %c%==1 (
			set satr=%satr%��
		) else (
			if %c%==2 (
				set satr=%satr%��
			) else (
				if %c%==3 (
					set satr=%satr%��
				) else (
					if %c%==4 (
						set satr=%satr%���
					) else (
						if %c%==5 (
							set satr=%satr%����
						) else (
							if %c%==6 (
								set satr=%satr%���
							) else (
								if %c%==7 (
									set satr=%satr%����
								) else (
									if %c%==8 (
										set satr=%satr%����
									) else (
										if %c%==9 (
											set satr=%satr%���
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