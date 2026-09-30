<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\MahasiswaResource;
use App\Models\Mahasiswa;
use App\Models\ProgramStudi;
use Illuminate\Http\Request;

class ProgramStudiController extends Controller
{
    public function mahasiswa(Request $request, ProgramStudi $programStudi)
    {
        $perHalaman = min($request->integer('per_halaman', 10), 100);

        $daftar = Mahasiswa::where('program_studi_id', $programStudi->id)
            ->with('programStudi')
            ->orderBy('nama')
            ->paginate($perHalaman);

        return MahasiswaResource::collection($daftar);
    }
}