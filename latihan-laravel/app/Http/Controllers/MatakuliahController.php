<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $daftarMatakuliah = [
            ['kode' => 'TKA101', 'nama' => 'Pemrograman Dasar', 'sks' => 3],
            ['kode' => 'TKA102', 'nama' => 'Sistem Operasi', 'sks' => 3],
            ['kode' => 'TKA103', 'nama' => 'Pendidikan Pancasila', 'sks' => 2],
            ['kode' => 'TKA104', 'nama' => 'Matematika Diskrit', 'sks' => 3],
            ['kode' => 'TKA105', 'nama' => 'Bahasa Inggris', 'sks' => 2],
            
        ];

        $katakunci = $request->query('q', '');
        if ($katakunci !== '') {
            $hasilPencarian = [];
            foreach ($daftarMatakuliah as $mk) {
                if (stripos($mk['nama'], $katakunci) !== false || stripos($mk['kode'], $katakunci) !== false) {
                    $hasilPencarian[] = $mk;
                }
            }
            $daftarMatakuliah = $hasilPencarian;
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $daftarMatakuliah,
            'katakunci' => $katakunci
        ]);
    }

    public function show(string $kode)
    {
        return view('matakuliah.show', ['kode' => $kode]);
    }

    public function cari(Request $request)
    {
        $katakunci = $request->query('q', '');
        return response()->json([
            'kata_kunci' => $katakunci,
            'metode' => $request->method(),
            'path' => $request->path(),
        ]);
    }

}