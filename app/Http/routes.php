<?php

Route::get('/', function()
{
	return view('home');
});

Route::get('/login', function()
{
	return view('login');
});

Route::post('/login', function()
{
	if(Auth::attempt([
		'email' => Input::get('email'),
		'password' => Input::get('password'),
	])) {
		return redirect()->intended('/');
	}
	return redirect()->back();
});

Route::get('/logout', function()
{
	Auth::logout();
	return redirect('/');
});

Route::get('/articles', function()
{
	if(Auth::check()) {
		$articles = App\Article::all();
	} else {
		$articles = App\Article::where('is_published', true)->get();
	}
	return view('articles.list', [
		'articles' => $articles,
	]);
});

Route::get('/article/{article}', function(App\Article $article)
{
	return view('articles.show', [
		'article' => $article,
	]);
});

Route::group(['middleware' => 'auth'], function()
{
	Route::get('/articles/new', function()
	{
		return view('articles.new');
	});
	
	Route::post('/articles/new', function(App\Http\Requests\NewArticleRequest $request)
	{
		$article = App\Article::create($request->all());
		return redirect('/article/'.$article->slug);
	});
	
	Route::get('/articles/{article}/publish', function(App\Article $article)
	{
		$article->is_published = true;
		$article->save();
		return redirect()->back();
	});
	
	Route::get('/articles/{article}/unpublish', function(App\Article $article)
	{
		$article->is_published = false;
		$article->save();
		return redirect()->back();
	});
	
	Route::get('/articles/{article}/delete', function(App\Article $article)
	{
		$article->delete();
		return redirect()->back();
	});
});
