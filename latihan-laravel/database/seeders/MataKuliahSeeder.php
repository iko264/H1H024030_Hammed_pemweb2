<?php

namespace Database\Seeders;

use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class MataKuliahSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $daftar = [
            ['kode' => 'TK101', 'nama' => 'Pemrograman Web II', 'sks' => 4, 'semester' => 4],
            ['kode' => 'TK102', 'nama' => 'Struktur Data', 'sks' => 3, 'semester' => 2],
            ['kode' => 'TK103', 'nama' => 'Basis Data', 'sks' => 3, 'semester' => 3],
        ];

        foreach ($daftar as $item) {
            Matakuliah::create($item);
        }
    }
}
