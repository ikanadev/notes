## Django commands

Be sure to have everything [installed](https://docs.djangoproject.com/en/3.1/intro/install/) correctly.
Or you can check if your linux package manager has the next packages: `python3` `python3-pip` `python3-django`, if that is the case, you have everything to continue :D

A collection of most common django commands

- `django-admin startproject myproject` Creates a new project boilerplate.
- `python manage.py runserver` Run the development server.
- `python manage.py startapp myapp` Creates a new app boilerplate. An app is a program which does one task and can be used in many projects, a project can have many apps.
- `python manage.py migrate` Look the installed apps and creates all the database tables needed.
- `python manage.py makemigrations myapp` When you make changes in your models, you need to specify the app.
- `python manage.py sqlmigrate polls 000X` Show the sql which will be used to make the migrations, you need the specify the number of the migration (usually 000X)
- `python manage.py shell` Loads the python REPL shell where you can play with the models.
- `python manage.py createsuperuser` Creates a new super user for the admin panel.

