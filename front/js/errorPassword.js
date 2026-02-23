document.addEventListener('DOMContentLoaded', function() {
    const form = document.querySelector("form");
    const pass1 = document.getElementById("password");
    const pass2 = document.getElementById("password2");
    const errorBox = document.getElementById("errorMessage");
    const submitBtn = document.querySelector(".btn");
    let timeoutId;

    // Добавляем элементы для анимации отправки
    const progressBar = document.createElement('div');
    progressBar.className = 'progress-bar';
    progressBar.id = 'progressBar';
    progressBar.innerHTML = '<div class="progress-bar-fill"></div>';
    form.appendChild(progressBar);

    const successMessage = document.createElement('div');
    successMessage.className = 'success-message';
    successMessage.id = 'successMessage';
    successMessage.textContent = 'Регистрация прошла успешно! Перенаправляем...';
    form.appendChild(successMessage);

    function showError(message) {
        if (timeoutId) {
            clearTimeout(timeoutId);
        }

        errorBox.textContent = message;
        errorBox.classList.add("active");

        pass1.classList.add("error");
        pass2.classList.add("error");

        timeoutId = setTimeout(() => {
            errorBox.classList.remove("active");
            pass1.classList.remove("error");
            pass2.classList.remove("error");
        }, 3000);
    }

    function validateOnInput() {
        if (pass2.value.length > 0) {
            if (pass1.value !== pass2.value) {
                showError("Пароли не совпадают");
            } else {
                errorBox.classList.remove("active");
                pass1.classList.remove("error");
                pass2.classList.remove("error");
            }
        }
    }

    pass1.addEventListener('input', validateOnInput);
    pass2.addEventListener('input', validateOnInput);

    form.addEventListener("submit", async function(e) {
        // Проверка длины пароля
        if (pass1.value.length < 6) {
            e.preventDefault();
            showError("Пароль должен быть минимум 6 символов");
            pass1.focus();
            return;
        }

        // Проверка совпадения паролей
        if (pass1.value !== pass2.value) {
            e.preventDefault();
            showError("Пароли не совпадают");
            pass1.value = "";
            pass2.value = "";
            pass1.focus();
            return;
        }

        // Если все проверки пройдены - показываем анимацию отправки
        e.preventDefault(); // Отменяем стандартную отправку для анимации

        // Показываем анимацию загрузки
        submitBtn.classList.add('loading');
        form.classList.add('form-sending');
        progressBar.classList.add('active');

        // Скрываем сообщение об ошибке если было
        errorBox.classList.remove('active');

        try {
            // Собираем данные формы
            const formData = new FormData(form);
            const data = {
                username: formData.get('username'),
                email: formData.get('email'),
                password: formData.get('password')
            };

            // Отправляем данные на сервер
            const response = await fetch('/registr', {
                method: 'POST',

                body: new FormData(form)
            });

            // Имитация задержки (можно убрать в продакшене)
            await new Promise(resolve => setTimeout(resolve, 1500));

            if (response.ok) {
                // Успешная отправка
                submitBtn.classList.remove('loading');
                submitBtn.classList.add('success');
                submitBtn.textContent = '✓ Успешно!';

                progressBar.classList.remove('active');
                successMessage.classList.add('active');

                // Очищаем форму
                form.reset();

                // Перенаправление через 2 секунды
                setTimeout(() => {
                    window.location.href = '/success';
                }, 2000);
            } else {
                throw new Error('Ошибка сервера');
            }
        } catch (error) {
            console.error('Error:', error);

            // Показываем ошибку
            submitBtn.classList.remove('loading');
            submitBtn.classList.add('error');
            submitBtn.textContent = 'Ошибка!';

            progressBar.classList.remove('active');
            form.classList.remove('form-sending');

            showError('Произошла ошибка при отправке. Попробуйте снова.');

            // Возвращаем кнопку в исходное состояние через 2 секунды
            setTimeout(() => {
                submitBtn.classList.remove('error');
                submitBtn.textContent = 'Зарегистрироваться';
            }, 2000);
        }
    });
});