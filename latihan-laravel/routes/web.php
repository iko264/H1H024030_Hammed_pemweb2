<?php

//use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MahasiswaController;
use App\Http\Controllers\MatakuliahController;

Route::get('/data-mahasiswa', [MahasiswaController::class,
'index'])->name('mahasiswa.index');

Route::get('/data-mahasiswa/{nim}', [MahasiswaController::class,
'show'])->name('mahasiswa.show');

Route::get('/cari-mahasiswa', [MahasiswaController::class, 'cari']);

Route::get('/matakuliah', [MatakuliahController::class, 'index'])->name('matakuliah.index');
Route::get('/matakuliah/{kode}', [MatakuliahController::class, 'show'])->name('matakuliah.show');
Route::get('/cari-matakuliah', [MatakuliahController::class, 'cari']);
